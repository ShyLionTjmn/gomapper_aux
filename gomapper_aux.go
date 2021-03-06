package gomapper_aux

import (
  "strconv"
  "strings"
  "errors"
)

func LastResultDecode(s string) (string, int64, int64, string, error) {
  a:= strings.Split(s, ":")
  if len(a) < 4 { return "", 0, 0, "", errors.New("bad last_result format") }

  i1, err := strconv.ParseInt(a[1], 10, 64)
  if err != nil { return "", 0, 0, "", errors.New("bad last_result format") }

  i2, err := strconv.ParseInt(a[2], 10, 64)
  if err != nil { return "", 0, 0, "", errors.New("bad last_result format") }

  return a[0], i1, i2, strings.Join(a[3:], ":"), nil
}

func SafeDevId(s string) string {
  ret := s
  ret = strings.ReplaceAll(ret, " ", "_")
  ret = strings.ReplaceAll(ret, "/", "s")
  ret = strings.ReplaceAll(ret, ":", "c")
  ret = strings.ReplaceAll(ret, "\t", "_")
  ret = strings.ReplaceAll(ret, ">", "_")
  ret = strings.ReplaceAll(ret, "<", "_")
  return ret
}

func SafeIntId(s string) string {
  ret := s
  ret = strings.ReplaceAll(ret, " ", "_")
  ret = strings.ReplaceAll(ret, "\t", "_")
  ret = strings.ReplaceAll(ret, ">", "_")
  ret = strings.ReplaceAll(ret, "<", "_")
  ret = strings.ReplaceAll(ret, "/", "s")
  ret = strings.ReplaceAll(ret, ":", "c")
  return ret
}
