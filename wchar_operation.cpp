int wide_strlen(const wchar_t *str1)
{
  int len = 0;
  while (str1[ len ])
  {
    len ++;
  }
  return len;
}

char *unicode_to_ansi_copy(char * dest, int dest_len,
                           const wchar_t *src, size_t buffer_len)
{
  if (!src || !dest)
  {
    return NULL;
  }

  buffer_len = wide_strlen(src) + 1;

  int i;
  for (i = 0; i < buffer_len && i < dest_len && src[ i ] != 0; i ++)
  {
    dest[ i ] = src[ i ] & 0x00FF;
  }

  dest[ i ] = '\0';
  return dest;
}
