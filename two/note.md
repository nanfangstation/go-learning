每个可执行go程序特征：
- 声明main函数作为程序入口
- 包名main

## 程序架构
![](https://tva1.sinaimg.cn/large/0081Kckwgy1glshcnt0grj317g0u0ng8.jpg)

## 项目结构
```
- sample
    - data
        data.json -- 包含一组数据源
    - matchers
        rss.go -- 搜素rss源的匹配器
    - search
        default.go -- 搜索数据用的默认匹配器
        feed.go -- 用于读取数据文件
        match.go --用于支持不同匹配器的接口
        search.go -- 执行搜索的主控制逻辑
    main.go -- 程序入口
```