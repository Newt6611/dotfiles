vim.pack.add({
  { src = 'https://github.com/folke/trouble.nvim' }
})

require('trouble').setup()

local map = vim.keymap.set
map('n', '<leader>xx', '<cmd>Trouble diagnostics toggle<CR>', { desc = 'Diagnostics (Trouble)', })
