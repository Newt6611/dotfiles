return {
  'stevearc/oil.nvim',
  opts = {
    use_default_keymaps = false,
    keymaps = {
        ['<CR>'] = 'actions.select',
        ['-'] = { 'actions.parent', mode = 'n' },
        ['_'] = { 'actions.open_cwd', mode = 'n' },
    },
    view_options = {
      show_hidden = true,
    },
    float = {
      border = 'rounded',
    }
  },
  -- dependencies = { 'nvim-tree/nvim-web-devicons' }, -- use if you prefer nvim-web-devicons
  -- Lazy loading is not recommended because it is very tricky to make it work correctly in all situations.
  lazy = false,
  keys = {
    { '-', mode = { 'n' }, function() require('oil').open_float() end, desc = 'Flash Treesitter' },
  }
}
