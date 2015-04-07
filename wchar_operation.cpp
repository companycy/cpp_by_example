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

std::string ws2s(const std::wstring& ws)
{
  std::string curLocale = setlocale(LC_ALL, NULL);        // curLocale
                                                          // = "C";
  setlocale(LC_ALL, "chs");
  const wchar_t* _Source = ws.c_str();
  size_t _Dsize = 2 * ws.size() + 1;
  char *_Dest = new char[_Dsize];
  memset(_Dest,0,_Dsize);
  wcstombs(_Dest,_Source,_Dsize);
  std::string result = _Dest;
  delete []_Dest;
  setlocale(LC_ALL, curLocale.c_str());
  return result;
}


std::wstring s2ws(const std::string& s)
{
  setlocale(LC_ALL, "chs");
  const char* _Source = s.c_str();
  size_t _Dsize = s.size() + 1;
  wchar_t *_Dest = new wchar_t[_Dsize];
  wmemset(_Dest, 0, _Dsize);
  mbstowcs(_Dest,_Source,_Dsize);
  std::wstring result = _Dest;
  delete []_Dest;
  setlocale(LC_ALL, "C");
  return result;
}


void Wchar_tToChar(const wchar_t *wchar, char *dst, size_t len)
{
  DWORD dwNum = WideCharToMultiByte(CP_OEMCP,NULL,wchar,-1,NULL,0,NULL,FALSE);// WideCharToMultiByte
  WideCharToMultiByte (CP_OEMCP,NULL,wchar,-1,dst,dwNum,NULL,FALSE); //
}




