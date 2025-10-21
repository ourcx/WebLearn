---
# 首页元数据配置
layout: home
hero:
  name: FEI Design
  text: 企业级中后台 UI 解决方案
  tagline: 高效 | 灵活 | 优雅
  image:
    src: https://s2.loli.net/2025/06/29/wiHWDFjzmEVYCT3.jpg
    alt: FEI
  actions:
    - theme: brand
      text: 快速开始 →
      link: ./guide/getting-started.md
    - theme: alt
      text: 组件预览
      link: /components/button
features:
  - title: 开箱即用
    icon: 🚀
    details: 预设企业级标准配置，零配置开箱即用，支持 TypeScript 类型推导
  - title: 灵活定制
    icon: 🎨
    details: CSS 变量深度主题定制，支持动态切换主题与组件级样式覆盖
  - title: 性能优化
    icon: ⚡
    details: Tree-shaking 按需加载，<b>gzip</b> 后体积仅 <b>28KB</b>，组件懒加载支持
  - title: 进阶特性
    icon: 🔮
    details: 支持暗黑模式、国际化、无障碍访问(A11y)，提供完整 TypeScript 类型定义
---

<!-- 快速开始代码块 -->
::: code-group#quick-start

```bash [安装]
# NPM
npm install fei-design@next

# Yarn
yarn add fei-design@next

# PNPM
pnpm add fei-design@next