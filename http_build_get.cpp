// host: www.baidu.com
// page: /
// thus result looks like:
// GET / HTTP/1.0
// Host: www.baidu.com
// User-Agent: HTMLGET 1.0
void build_get_query(const char *host, const char *page,
                     char *query, size_t len) {
  const char *tpl =
    "GET %s HTTP/1.0\r\nHost: %s\r\nUser-Agent: %s\r\n\r\n"; // 1.0 version
  // "GET /%s HTTP/1.1\r\nHost: %s\r\nUser-Agent: %s\r\nConnection: close\r\n\r\n";

  // -5 is to consider the %s %s %s in tpl and the ending \0
  // char *query =
  //   (char *)malloc(strlen(host)+strlen(page)+strlen(USERAGENT)+strlen(tpl)-5);
  // sprintf(query, tpl, page, host, USERAGENT);
  _snprintf(query, len, tpl, page, host, USERAGENT);
}
