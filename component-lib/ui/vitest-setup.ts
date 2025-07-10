// vitest-setup.ts
import { config } from '@vue/test-utils';
import { JSDOM } from 'jsdom';

// 创建完整的 DOM 环境
const dom = new JSDOM('<!DOCTYPE html><html><body></body></html>', {
  url: 'http://localhost/'
});

global.window = dom.window as any;
global.document = dom.window.document;
global.navigator = dom.window.navigator;
global.HTMLElement = dom.window.HTMLElement;
global.Node = dom.window.Node;

// 配置 Vue Test Utils
config.global = {
  ...config.global,
  stubs: {
    // 全局存根
    Transition: true,
    'transition-group': true
  },
  mocks: {
    // 全局模拟
  }
};

// 添加 Vue 3 运行时所需的全局方法
global.requestAnimationFrame = (callback) => setTimeout(callback, 0);
global.cancelAnimationFrame = (id) => clearTimeout(id);