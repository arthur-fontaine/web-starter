const path = require("node:path");
const { mergeConfig } = require('metro-config');
const { getDefaultConfig } = require('expo/metro-config');
const packageJson = require('./package.json');

/** @type {Object.<string, string>} */
const extraNodeModules = {};
for (const [key, value] of Object.entries(packageJson.dependencies)) {
  if (!value.startsWith('file:')) continue;
  const dependencyPath = value.replace(/^file:/, '');
  extraNodeModules[key] = path.resolve(__dirname, dependencyPath);
}

const config = {
  watchFolders: Object.values(extraNodeModules),
  resolver: {
    extraNodeModules: new Proxy(extraNodeModules, {
      get: (target, name) =>
        // @ts-ignore
        name in target ? target[name] : path.join(process.cwd(), `node_modules/${name.toString()}`),
    }),
    unstable_enablePackageExports: true,
  },
  resetCache: true,
};

module.exports = mergeConfig(getDefaultConfig(__dirname), config);
