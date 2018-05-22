package main

var (
	actionToString = map[string]map[string]string{

		"todolists": map[string]string{
			"created":    "[%s] “%s” 创建了任务清单 “%s” ",
			"deleted":    "[%s] “%s” 删除了任务清单 “%s” ",
			"archived":   "[%s] “%s” 归档了任务清单 “%s” ",
			"unarchived": "[%s] “%s” 重新激活了任务清单 “%s” ",
			"updated":    "[%s] “%s” 更新了任务清单 “%s” ",
			"commented":  "[%s] “%s” 评论了任务清单 “%s” ： “%s” ",
		},
		"todos": map[string]string{
			"created":          "[%s] [%s] “%s” 创建了任务 “%s” ",
			"updated":          "[%s] [%s] “%s” 更新了任务 “%s” ",
			"assigned":         "[%s] [%s] “%s” 把任务 “%s” 指派给了 “%s” ",
			"unassigned":       "[%s] [%s] “%s” 取消了 “%s” 的任务 “%s” ",
			"deadline_changed": "[%s] [%s] “%s” 将 “%s” 的截止时间设为 “%s” ",
			"moved":            "[%s] [%s] “%s” 把任务 “%s” 移动到 “%s” 中",
			"started":          "[%s] [%s] “%s” 开始处理这条任务 “%s” ",
			"paused":           "[%s] [%s] “%s” 暂停处理这条任务 “%s” ",
			"completed":        "[%s] [%s] “%s” 完成了这条任务 “%s” ",
			"deleted":          "[%s] [%s] “%s” 删除了这条任务 “%s” ",
			"reopened":         "[%s] [%s] “%s” 重新打开了这条任务 “%s” ",
			"commented":        "[%s] [%s] “%s” 评论了任务 “%s” ： “%s” ",
		},
		"check_items": map[string]string{
			"created":          "[%s] “%s” 在任务 “%s” 创建了检查清单 “%s” ",
			"updated":          "[%s] “%s” 在任务 “%s” 更新了检查清单 “%s” ",
			"assigned":         "[%s] “%s” 在任务 “%s” 把检查清单 “%s” 指派给了 “%s” ",
			"unassigned":       "[%s] “%s” 在任务 “%s” 取消了 “%s” 的检查清单 “%s” ",
			"deadline_changed": "[%s] “%s” 在任务 “%s” 将 “%s” 的截止时间设为 “%s” ",
			"completed":        "[%s] “%s” 在任务 “%s” 完成了检查清单 “%s” ",
			"deleted":          "[%s] “%s” 在任务 “%s” 删除了检查清单 “%s” ",
			"reopened":         "[%s] “%s” 在任务 “%s” 重新打开了检查清单 “%s” ",
		},

		"topics": map[string]string{
			"created":    "[%s] “%s” 创建了讨论 “%s” ",
			"deleted":    "[%s] “%s” 删除了讨论 “%s” ",
			"sticked":    "[%s] “%s” 置顶了讨论 “%s” ",
			"unsticked":  "[%s] “%s” 取消置顶了讨论 “%s” ",
			"archived":   "[%s] “%s” 结束了讨论 “%s” ",
			"unarchived": "[%s] “%s” 重新打开了讨论 “%s” ",
			"commented":  "[%s] “%s” 评论了讨论 “%s” ： “%s” ",
		},

		"documents": map[string]string{
			"created":   "[%s] “%s” 创建了文档 “%s” ",
			"deleted":   "[%s] “%s” 删除了文档 “%s” ",
			"updated":   "[%s] “%s” 编辑了文档 “%s” ",
			"recovered": "[%s] “%s” 恢复了文档 “%s” ",
			"commented": "[%s] “%s” 评论了文档 “%s” ： “%s” ",
		},

		"attachments": map[string]string{
			"created":   "[%s] “%s” 上传了文件 “%s” ",
			"deleted":   "[%s] “%s” 删除了文件 “%s” ",
			"updated":   "[%s] “%s” 更新了文件 “%s” ",
			"commented": "[%s] “%s” 评论了文件 “%s” ： “%s” ",
		},
	}
)
