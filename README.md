# ApkUrlGrep
Tool that allow extract endpoints from APK files
![alt text](https://i.ibb.co/V3nFRwJ/image-2020-05-06-21-30-30.png)

## Install
1) Install `apkurlgrep`
```
▶ go get -u github.com/ndelphit/apkurlgrep
```
2) Install [apktool](https://ibotpeaches.github.io/Apktool/install/)


## Usage


```
▶ apkurlgrep -a ~/path/to/file.apk
Result of URLs:

https://example.com
https://example.net
https://example.edu

Result of URLs Paths:

/example
/admin
/onboarding
```
