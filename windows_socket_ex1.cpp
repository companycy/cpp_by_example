
void post_msg(const char *host, const char *page,
              const char *buf, const size_t len) {
  WSADATA  Ws;
  if (WSAStartup(MAKEWORD(2,2), &Ws) != 0) {
    printf("Init Windows Socket Failed::%s\n", GetLastError());
    return;
  }

  SOCKET ClientSocket = socket(AF_INET, SOCK_STREAM, IPPROTO_TCP);
  if (ClientSocket == INVALID_SOCKET) {
    printf("Create Socket Failed::%s\n", GetLastError());
    return;
  } else {
    printf("Created Socket!\n");
  }

  struct sockaddr_in ServerAddr;
  ServerAddr.sin_family = AF_INET;
  ServerAddr.sin_addr.s_addr = inet_addr(host);
  ServerAddr.sin_port = htons(80);
  memset(ServerAddr.sin_zero, 0x00, 8);

  int Ret =
    connect(ClientSocket,(struct sockaddr*)&ServerAddr, sizeof(ServerAddr));
  if (Ret == SOCKET_ERROR) {
    printf("Connect Error::%s\n", GetLastError());
    return;
  } else {
    printf("Connected Socket!\n");
  }

  // const size_t SendBufferLen = 4096; // 2 * len;
  // char *SendBuffer = new char[SendBufferLen]; // = {0};
  char SendBuffer[4096] = {0};
  build_post_query(host, page, buf, SendBuffer, sizeof(SendBuffer));
  if (strstr(SendBuffer, "hr_hi_person") && strstr(SendBuffer, "INSERT")) {
    fprintf(stdout, "insert into table info\n");
  }
  fprintf(stdout, "Query <<START>>\n%sQuery <<END>>\n", SendBuffer);

  Ret = send(ClientSocket, SendBuffer, (int)strlen(SendBuffer), 0);
  if (Ret == SOCKET_ERROR) {
    printf("Send Info Error::%s\n", GetLastError());
    return;
  } else {
    printf("Sent Info!\n");
  }
  // delete []SendBuffer;



  closesocket(ClientSocket);
  WSACleanup();
