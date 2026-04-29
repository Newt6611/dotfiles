return {
  'lewis6991/gitsigns.nvim',
  lazy = false,
  opts = {
    current_line_blame = true,
    current_line_blame_opts = {
      delay = 0,
    },
  },
  keys = {
    {
      '<leader>r',
      mode = { 'n', 'v' },
      function()
        require('gitsigns').reset_hunk()
      end,
      desc = 'Reset Hunk',
    },
    {
      '<leader>b',
      mode = { 'n' },
      function()
        require('gitsigns').blame_line()
      end,
      desc = 'Blam line'
    },
    {
      '<leader>j',
      mode = { 'n' },
      function()
        require('gitsigns').nav_hunk('next')
      end,
      desc = 'Next hunk'
    },
    {
      '<leader>k',
      mode = { 'n' },
      function()
        require('gitsigns').nav_hunk('prev')
      end,
      desc = 'Prev hunk'
    },
  },
}
