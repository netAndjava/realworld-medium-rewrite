# go real world design

1. 用户(user)注册 medium，用户登录 medium,用户退出登录,用户修改个人信息(profile)

2. 用户在使用 medium 之前选择主题(topics)，这样可以给用户推荐相关主题出版社的文章和独立作者的文章,用户可以在个人中心修改自己感兴趣的(topics)

3. 用户可以写文章，把文章保存为草稿(drafts), 给文章添加标签(tags)然后发布文章。可以在个人中心查看和修改文章,也可以删除文章.

4. 给文章生成友好链接(friend link)供朋友免费查看

5. 用户可以创建出版社(publication), 给出版社添加标签(tags)，邀请用户成为出版社的编辑(editor)来帮助管理出版社，邀请作者向出版社投稿，如果作者加入出版社，可以将文章投稿给出版社。用户可以在个人中心中管理出版社

6. 作者可以向加入的出版社投稿，也可以从出版社中删除自己投稿的文章

7. 出版社编辑可以浏览作者提交的文章，添加文章，修改和发布提交的文章，当觉得文章不适合出版社时也可以删除文章

8. 用户可以浏览不同出版社的文章列表，如果对文章感兴趣可以收藏(collect),不感兴趣可以可以忽略(dismiss)这篇文章,收藏的文章列表可以在个人中心查看。

9. 用户在浏览各个出版社的文章时，如果对出版社的文章比较感兴趣可以关注出版社，这样可以收到出版社后续发版的文章，如果不感兴趣可以屏蔽(mute)出版社,mute 与 unfollow 的关系

10. 浏览文章时，对于感兴趣的文章，可以关注(follow)文章作者来知道作者后续的创作，如果不想再看该作者的相关文章可以屏蔽该作者(mute author)。

11. 用户阅读文章时，可以把文章分享(share)给朋友，如果觉得文章写的好可以给文章点赞(like)，如果想跟作者或其他人交流探讨文章，可以在文章下面评论(comment)文章,作者可以回复你的评论

12. 用户可以查看自己收藏的文章，如果不感兴趣了可以取消收藏

13. medium 会记录用户的浏览历史(history),用户可以查看最近浏览的历史

模拟 medium.com;

medium 提供了如下`部分`功能,(有些部分是我在思考中推导的):

1. 让用户(user)可以在 medium 写文章(article);
2. 你也可以在 medium 上浏览(find[search],view)其他用户写的文章(article);
3. 当你在看其他作者文章的时候:你如果喜欢这篇文章,你可以给这篇文章点赞(like)和收藏(collect)到你的收藏夹(Favorites)中;
4. 如果你觉得这个作者写的东西不错,可能会想知道这个作者后续创作的东西能自动通知你,你可以关注(flow)这个作者.
5. 由于你的观点和喜好可能会有相同喜好的人,所以我们可能把你的喜好和关注也透明(notify)给其他关注(flow)你的人.
6. 当你在读其他作者的文章的时候可能会促进你的创作欲望,你可以根据这篇文章进行二次创作(Innovation),原文章可以在你的文章中表现为引用(毕竟你参考了他),你的这次创作可能被通知到你引用的作者,因为可能你和这个作者都在同一个领域问题思考,原作者可能会过来和你一起讨论或下一步的(共同创作).也有可能觉得你弱爆了,无视你.

7. 用户在创作的过程中是如何创作的,这个不是一个简单的写文章,而是关注用户在知识学习,收集,讨论,整理,创作,验证,的一个 infinity 过程.

8. 社交关系的处理要为影响或协同创作服务或产生有(当前社会)创始人定位具备价值意义的服务方式,避免出现社交网络中的键政,无意义的吐槽,攻击,因此社交行为会具有惩罚性措施;
9. 产品的定位是关注用户终身学习和帮助提供用户生活行为习惯.
10. 创作不再是简单的独立一篇博文,而是如何帮助用户学习,思考,创作,找到有志同道合的人共同成长和创作.
