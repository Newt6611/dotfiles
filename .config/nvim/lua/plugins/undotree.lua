vim.pack.add({
  { src = 'https://github.com/jiaoshijie/undotree' }
})

local undotree = require('undotree')
undotree.setup()

local map = vim.keymap.set
map('n', '<leader>ud', undotree.toggle, { desc = 'Toggle undotree' })
