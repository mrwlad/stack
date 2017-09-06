import chalk from 'chalk';

export const log = txt => {
  process.stdout.write(`${txt}\n`);
};

export const err = txt => {
  process.stderr.write(`\n${chalk`{red ERROR: ${txt}}`}\n`);
};

export const warn = txt => {
  process.stderr.write(`\n${chalk`{yellow WARNING: ${txt}}`}\n`);
};
