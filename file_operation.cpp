

void write_log2(const int val)
{
  FILE *fp = fopen("c://oledb.log", "ab+");
  if (!fp)
    return;

  fprintf(fp, "value: %d\n", val);
  fclose(fp);
}
