//go:build ignore
#include <stdio.h>
#include <netdb.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <stdlib.h>
#include <string.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <unistd.h>
#include <signal.h>

#define MAXMSGSIZE 1024
#define PORT 8080

int sockfd;

void chat(int connfd)
{
	char buffer[MAXMSGSIZE];
	int n;
	while (1) {
        n = recv(connfd, buffer, MAXMSGSIZE, MSG_WAITFORONE);
        buffer[n] = '\0';

		printf("Client: %s\n", buffer);
		send(connfd, buffer, n, MSG_CONFIRM);

		if (strncmp("/exit", buffer, 5) == 0) {
			printf("Server exit...\n");
			close(connfd);
            exit(0);
		}
	}
}

void sig_handler(int signal) {
    printf("Termination...\n");
	close(sockfd);
    exit(0);
}

int main()
{
	int connfd, len;
	struct sockaddr_in servaddr, cliaddr;

	sockfd = socket(AF_INET, SOCK_STREAM, 0);
	if (sockfd == -1) {
		printf("Socket creation failed...\n");
		return -1;
	} else {
		printf("Socket successfully created..\n");
    }
	memset(&servaddr, 0, sizeof(servaddr));

	servaddr.sin_family = AF_INET;
	servaddr.sin_addr.s_addr = htonl(INADDR_ANY);
	servaddr.sin_port = htons(PORT);

	if ((bind(sockfd, (struct sockaddr*)&servaddr, sizeof(servaddr))) != 0) {
		printf("Socket bind failed...\n");
		return -1;
	} else {
		printf("Socket successfully binded..\n");
    }

    signal(SIGTERM, &sig_handler);
    signal(SIGINT, &sig_handler);

    while (1) {
        if ((listen(sockfd, 5)) != 0) {
            printf("Listen failed...\n");
            return -1;
        } else {
            printf("Server listening..\n");
        }
        len = sizeof(cliaddr);

        connfd = accept(sockfd, (struct sockaddr*)&cliaddr, &len);
        if (connfd < 0) {
            printf("Server accept failed...\n");
            return -1;
        } else {
            printf("Server accept the client...\n");
        }
        printf("Client IP: %s:%d\n", inet_ntoa(cliaddr.sin_addr), cliaddr.sin_port);

        int child = fork();
        if (child == 0) {
            chat(connfd);
        } else if (child == -1) {
            printf("Error with fork\n");
        }
    }
}