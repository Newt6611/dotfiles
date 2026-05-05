return {
  'folke/snacks.nvim',
  opts = {
    picker = {
      enabled = true,
      win = {
        -- input window
        input = {
          keys = {
            ["K"] = { "preview_scroll_up", mode = { "i", "n" } },
            ["J"] = { "preview_scroll_down", mode = { "i", "n" } },
            ["L"] = { "preview_scroll_right", mode = { "i", "n" } },
            ["H"] = { "preview_scroll_left", mode = { "i", "n" } },
          },
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
    -- {
    --   '<leader>gl',
    --   function() require('snacks').lazygit.log() end,
    --   desc = 'LazyGit Log',
    -- },
  },
}
