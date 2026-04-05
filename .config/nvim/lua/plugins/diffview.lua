return {
  'sindrets/diffview.nvim',
  lazy = true,
  keys = {
    {
      'df',
      mode = 'n',
      '<CMD>DiffviewOpen<CR>',
      desc = 'Open Diffview',
    },
    {
      'dc',
      mode = 'n',
      '<CMD>DiffviewClose<CR>',
      desc = 'Close Diffview',
    },
  }
}

