const { execSync } = require('child_process');

module.exports = {
  templatesOptions: {
    'new-day': {
      hooks: {
        postTemplateGeneration: (ctx) => {
          execSync(`go mod init aoc/day-${ctx.parametersValues.day} && go mod tidy`, {
            cwd: `${ctx.targetRoot}/day-${ctx.parametersValues.day}`,
            stdio: 'inherit',
          });
        },
      },
    },
  },
};
