// eslint-disable-next-line @typescript-eslint/no-require-imports
const path = require('path')

const buildEslintCommand = (filenames) => {
  return `next lint --max-warnings=0 --fix --file ${filenames
    .map((f) => path.relative(process.cwd(), f))
    .filter((f) => !f.startsWith('proto/'))
    .join(' --file ')}`
}

module.exports = {
  '*.{js,jsx,ts,tsx}': ['prettier --write', buildEslintCommand],
  '*.{json,css,md}': ['prettier --write'],
}
