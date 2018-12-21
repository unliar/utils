const fs = require('fs');
const path = require('path');

/**
 * 获取目标文件夹下所有的文件列表
 *
 * @param {string} targetPath 目标文件夹绝对路径
 * @param {Array} fileInfo 手动添加的文件配置信息
 */
const LoadAllFiles = (targetPath, fileInfo = []) => {
  const temp = [...fileInfo];
  const state = fs.statSync(targetPath);
  if (state.isFile()) {
    temp.push({filePath: targetPath});
    return temp;
  }
  if (state.isDirectory()) {
    const dirs = fs.readdirSync(targetPath);
    dirs.forEach(item => {
      const itemState = fs.statSync(path.join(targetPath, item));
      if (itemState.isFile()) {
        temp.push({filePath: path.join(targetPath, item)});
      }
      if (itemState.isDirectory()) {
        temp.push(
            ...LoadAllFiles(path.join(targetPath, item), fileInfo),
        );
      }
    });
    return temp;
  }
};

module.exports = LoadAllFiles;
