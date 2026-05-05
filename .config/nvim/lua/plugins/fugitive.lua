return {
  {
    'tpope/vim-fugitive',
    keys = {
      { '<leader>gg', '<cmd>Git<cr>', desc = 'Open Fugitive' },
    },
  },
  {
    'tpope/vim-rhubarb',
  },
  {
    "rbong/vim-flog",
    lazy = false,
    cmd = { "Flog", "Flogsplit", "Floggit" },
    dependencies = {
      "tpope/vim-fugitive",
    },
    keys = {
      { '<leader>gl', '<cmd>Flog<cr>', desc = 'Open Flog' },
    }
  },
}
