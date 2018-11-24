# golang-novel-spider
小说,go,spider
## 使用说明
 1. https://mp.weixin.qq.com/debug/cgi-bin/sandbox?t=sandbox/login 
 在wx/wx.go中配置自己申请的微信公众账号测试账号appid,secret,openid,模版消息id等
 
 2. 模版消息参考 小说: {{first.DATA}} 最新章节: {{keyword1.DATA}} {{remark.DATA}}
 
 3. 安装go的依赖包，go get 即可，需要 番****墙，移步蓝灯 https://github.com/getlantern/download
 
 4. 终端先设置set https_proxy=http://127.0.0.1:xxxxx, 再go get 
 
 5. 在main中替换filename和url，check方法中修改判断bookName 的方法
 
 6. Good Luck !
