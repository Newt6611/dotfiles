vim.pack.add({
  { src = 'https://github.com/saghen/blink.cmp' },

  -- optional: provides snippets for the snippet source
  { src = 'https://github.com/rafamadriz/friendly-snippets' }
})

require('blink.cmp').setup({
  keymap = {
    ['<Up>'] = { 'select_prev', 'fallback' },
    ['<Down>'] = { 'select_next', 'fallback' },
    ['<CR>'] = { 'accept', 'fallback' },
  }
})
