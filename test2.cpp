#include <cstdio>
#include <cstdlib>
#include <string>
#include <map>

using std::map;
using std::string;
using std::make_pair;

class Solution {
public:
  string largestNumber(vector<int> &num) {
    map<string, int> smap;
    for (vector<int>::iterator it = num.begin(); 
	 it != num.end(); it++) {
      char s[64] = {0};
      snprintf(s, sizeof(s), "%u", *it);
      smap.insert(make_pair(s, *it));
    }
        
    string result;
    for (map<string, int>::reverse_iterator it = smap.rbegin();
	 it != smap.end(); it++) {
      result += it->first;    
    }
    return result;
  }
};

int main() {
  const int l[] = {3, 30, 34, 5, 9};
  const int len = sizeof(l) / sizeof(int);
  map<string, int> smap;
  for (int i = 0; i < len; i++) {
    char s[32] = {0};
    snprintf(s, sizeof(s), "%u", l[i]);
    smap.insert(make_pair(s, l[i]));
  }

  string result;
  for (map<string, int>::reverse_iterator it = smap.rbegin();
       it != smap.rend(); it++) {
    result += it->first;
    printf("key->value: %s -> %d\n", it->first.c_str(), it->second);
  }

  fprintf(stdout, "%s\n", result.c_str());
  return 0;
}

