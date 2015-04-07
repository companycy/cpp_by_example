struct timeval
{
  long    tv_sec;         /* seconds */
  long    tv_usec;        /* and microseconds */
};

int gettimeofday(struct timeval* tv, void *)
{
  union
  {
    long long ns100;
    FILETIME ft;
  } now;
  GetSystemTimeAsFileTime (&now.ft);
  tv->tv_usec = (long) ((now.ns100 / 10LL) % 1000000LL);
  tv->tv_sec = (long) ((now.ns100 - 116444736000000000LL) / 10000000LL);

  return 0;
}

/*
          struct timeval tv_begin, tv_end, tv_diff;
          gettimeofday(&tv_begin, NULL);
          stat.timestamp =                // sql timestamp in milliseconds
            tv_begin.tv_sec * 1000 + tv_begin.tv_usec / 1000;

*/
