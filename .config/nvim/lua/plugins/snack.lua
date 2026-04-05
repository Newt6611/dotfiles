return {
  'folke/snacks.nvim',
  opts = {
    picker = {
      enabled = true,
    },
    indent = {
      enabled = true,
      animate = {
        enabled = false,
      },
    },
    lazygit = {
      enabled = true,
    },
  },
  keys = {
    {
      '<leader>pf',
      function() require('snacks').picker.files() end,
      desc = 'Find Files',
    },
    {
      '<leader>ps',
      function() require('snacks').picker.grep() end,
      desc = 'Grep',
    },
    {
      '<leader>gd',
      function() require('snacks').picker.git_diff() end,
      desc = 'Git Diff (Hunks)',
    },
    {
      '<leader>gl',
      function() require('snacks').lazygit.log() end,
      desc = 'LazyGit Log',
    },
    {
      '<leader>gg',
      function() require('snacks').lazygit() end,
      desc = 'LazyGit',
    },
  },
}
