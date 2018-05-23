package main

var (
	eventToString = map[string]string{
		"todolists":   "任务清单",
		"todos":       "任务",
		"check_items": "任务检查项目",
		"topics":      "讨论",
		"documents":   "文档",
		"attachments": "文件",
	}

	actionToString = map[string]string{
		"created":          "创建",
		"deleted":          "删除",
		"archived":         "归档",
		"unarchived":       "取消归档",
		"updated":          "更新",
		"commented":        "评论",
		"assigned":         "指派",
		"unassigned":       "取消指派",
		"deadline_changed": "设置截止时间",
		"moved":            "移动",
		"started":          "开始处理",
		"paused":           "暂停处理",
		"completed":        "完成",
		"reopened":         "重新打开",
		"sticked":          "置顶",
		"unsticked":        "取消置顶",
		"recovered":        "恢复",
	}

	eventFormat = map[string]map[string]string{

		"todolists": map[string]string{
			"created":    "创建了任务清单 “%s” ",
			"deleted":    "删除了任务清单 “%s” ",
			"archived":   "归档了任务清单 “%s” ",
			"unarchived": "重新激活了任务清单 “%s” ",
			"updated":    "更新了任务清单 “%s” ",
			"commented":  "评论了任务清单 “%s” ： “%s” ",
		},
		"todos": map[string]string{
			"created":          "创建了任务 “%s” ",
			"updated":          "更新了任务 “%s” ",
			"assigned":         "任务 “%s” 被指派给了 “%s” ",
			"unassigned":       "取消了 “%s” 的任务 “%s” ",
			"deadline_changed": "将 “%s” 的截止时间设为 “%s” ",
			"moved":            "任务 “%s” 被移动到 “%s” 中",
			"started":          "开始处理这条任务 “%s” ",
			"paused":           "暂停处理这条任务 “%s” ",
			"completed":        "完成了这条任务 “%s” ",
			"deleted":          "删除了这条任务 “%s” ",
			"reopened":         "重新打开了这条任务 “%s” ",
			"commented":        "评论了任务 “%s” ： “%s” ",
		},
		"check_items": map[string]string{
			"created":          "在任务 “%s” 里创建了检查清单 “%s” ",
			"updated":          "在任务 “%s” 里更新了检查清单 “%s” ",
			"assigned":         "在任务 “%s” 里把检查清单 “%s” 指派给了 “%s” ",
			"unassigned":       "在任务 “%s” 里取消了 “%s” 的检查清单 “%s” ",
			"deadline_changed": "在任务 “%s” 里将 “%s” 的截止时间设为 “%s” ",
			"completed":        "在任务 “%s” 里完成了检查清单 “%s” ",
			"deleted":          "在任务 “%s” 里删除了检查清单 “%s” ",
			"reopened":         "在任务 “%s” 里重新打开了检查清单 “%s” ",
		},

		"topics": map[string]string{
			"created":    "创建了讨论 “%s” ",
			"deleted":    "删除了讨论 “%s” ",
			"sticked":    "置顶了讨论 “%s” ",
			"unsticked":  "取消置顶了讨论 “%s” ",
			"archived":   "结束了讨论 “%s” ",
			"unarchived": "重新打开了讨论 “%s” ",
			"commented":  "评论了讨论 “%s” ： “%s” ",
		},

		"documents": map[string]string{
			"created":   "创建了文档 “%s” ",
			"deleted":   "删除了文档 “%s” ",
			"updated":   "编辑了文档 “%s” ",
			"recovered": "恢复了文档 “%s” ",
			"commented": "评论了文档 “%s” ： “%s” ",
		},

		"attachments": map[string]string{
			"created":   "上传了文件 “%s” ",
			"deleted":   "删除了文件 “%s” ",
			"updated":   "更新了文件 “%s” ",
			"commented": "评论了文件 “%s” ： “%s” ",
		},
	}
)
