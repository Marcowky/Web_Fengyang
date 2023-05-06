# 开发笔记

*用于开发中记录一些可以参考的要点*

# GIT

## GIT操作基本语句

以下是 Git 常用的基本命令行语句：

```bash
git init：初始化一个 Git 仓库。
git clone [url]：克隆（下载）一个 Git 仓库到本地。
git add [file]：将文件添加到暂存区。
git commit -m "message"：提交更改并添加提交消息。
git status：查看当前仓库的状态。
git log：查看提交历史。
git branch：列出所有本地分支，当前分支前面会标有一个 * 号。
git checkout [branch-name]：切换到指定分支。
git merge [branch]：将指定分支合并到当前分支。
git push：将本地代码推送到远程仓库。
git pull：从远程仓库拉取最新代码到本地。
git remote add [name] [url]：添加一个远程仓库。
git diff：查看修改前后的差异。
```

这些是 Git 基本的命令行语句，还有很多高级的 Git 命令可以使用，可以通过 git help 命令来查看 Git 的帮助文档。

## GIT版本控制

**文件状态**

在 Git 中，文件有以下几种状态：

> 未追踪（Untracked）：表示该文件不在 Git 的版本控制中。

> 已暂存（Staged）：表示该文件已被添加到 Git 的版本控制中，但还未提交到本地仓库。

> 已修改（Modified）：表示该文件已被修改但尚未被添加到暂存区。

> 已提交（Committed）：表示该文件已被添加到本地仓库，并且与远程仓库同步。

Git 可以通过 git status 命令来查看当前文件的状态。当文件发生变化时，可以使用 git add 将其添加到暂存区，然后使用 git commit 将其提交到本地仓库。

**流程**

创建文件->`Untracked`->git add->`Staged`->修改代码->`Modified`->git commit->`Committed`->git push->`代码上传至在线仓库`

**项目管理**


