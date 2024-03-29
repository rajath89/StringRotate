package main

import (
  "fmt"
  "bytes"
  "strings"
)

type rec string

func(str rec) printStr(){
  fmt.Printf("%s",str)
}

func(str rec) convertToBytes(){
  b:=[]byte(str)
  fmt.Println(b)
}

func(str rec) Rot1OnString() []byte{
  Rot1:= func(r rune) rune {
    switch {
      case r >= 'A' && r <= 'Z':
        return 'A' + (r-'A'+1)%26
      case r >= 'a' && r <= 'z':
        return 'a' + (r-'a'+1)%26
      }
      return r
  }
  var byted []byte
  byted=bytes.Map(Rot1, []byte(str))

  return byted
}



func(str rec) rotOnWholeString() ([]byte,int){
  total:=0
  //var by []byte
  var b bytes.Buffer
  ln:=len(str)
  b.Write(str.Rot1OnString())
  for total<ln{


    var rw rec

    tr1:=b.Next(ln)
    b.Reset()

    rw=rec(tr1)

    tr2:=rw.Rot1OnString()
    b.Write(tr2)


    total=total+1
  }
  tr:=b.Next(30)
  return tr,total
}




func main()  {
  var str1 rec="hello world"
  var stre []string
  stre=strings.Split(string(str1)," ")

  j:=  " "
  var b1 bytes.Buffer
  for strf := range stre{

    df,_:=rec(stre[strf]).rotOnWholeString()

    b1.Write(df)
    b1.Write([]byte(j))

}
  fmt.Println(string(b1.Next(len(str1))))
}
