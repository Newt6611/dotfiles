return {
  'lewis6991/gitsigns.nvim',
  lazy = false,
  opts = {},
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
  },
}
