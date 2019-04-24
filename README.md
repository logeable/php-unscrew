# php-unscrew

**PM9SCREW**

参考：[!php_screw的加密和解密探究](https://www.skactor.tk/2018/03/26/php-screw%E7%9A%84%E5%8A%A0%E5%AF%86%E5%92%8C%E8%A7%A3%E5%AF%86%E6%8E%A2%E7%A9%B6%EF%BC%88%E4%BA%8C%EF%BC%89%E8%A7%A3%E5%AF%86%E7%AE%97%E6%B3%95%E4%B8%8Epython%E5%AE%9E%E7%8E%B0/)

```
go build
./php-screw -f <file> -k <key>

objdump -s -j .data php_screw.so
第一行则为key(删除空白字符)
```
