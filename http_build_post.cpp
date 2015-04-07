void build_post_query(const char *host, const char *page, const char *buf,
                      char *query, size_t len) {
  const char *tpl =
    "POST %s HTTP/1.1\r\n"
    "Host: %s\r\n"
    "Content-Type: application/x-www-form-urlencoded\r\n"
    ""
    "Content-Length: %d\r\n\r\n"
    "%s";
  _snprintf(query, len, tpl, page, host, strlen(buf), buf);
}
