# Welcome to 智能桌宠! 🐾

本项目代码遵循 **Apache License 2.0** 开源协议。  
完整协议请查看仓库根目录中的 `LICENSE` 文件。

特别鸣谢：本项目基于 [Wails 3](https://wails.io/) 框架开发.

---

## 📂 项目结构

### 后端（backend）

所有后端业务逻辑都在 `backend` 文件夹中。  
请保持原有代码风格，我们采用 **类 DDD 架构**，整体应用通过依赖注入构建。
```text
backend/
├─ main server # 调度具体服务 server
├─ service server # 调度 controller，并执行前后处理逻辑
│ ├─ controller # 统筹 service 功能
│       ─ service # 操控 AI 等具体业务逻辑
```

### 前端（frontend）

- 待完善...
- 使用 Wails 框架渲染界面
- 计划支持桌面应用动画、交互与智能桌宠逻辑显示

---

## 💖 支持智能桌宠

如果你喜欢本项目，欢迎通过以下方式支持我们，让项目持续成长：

| 支持方式       | 链接 / 二维码                                                         |
|----------------|------------------------------------------------------------------|
| PayPal         | [https://www.paypal.me/hao97336](https://www.paypal.me/hao97336) |
| 支付宝 / 微信  | <img src="assets/pay.jpg" width="150"/>                           |

感谢你的支持！每一份赞助都能让智能桌宠更聪明、更可爱 🎉
