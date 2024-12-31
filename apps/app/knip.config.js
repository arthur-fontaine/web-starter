const packageJson = require('./package.json');

const softLinkedDependencies = Object.entries(packageJson.dependencies)
  .filter(([_, value]) => value.startsWith('file:'))
  .map(([key]) => key);

/** @type {import('knip').KnipConfig} */
module.exports = {
  ignoreDependencies: [
    ...softLinkedDependencies,
  ],
};