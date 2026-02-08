
local map = vim.keymap.set

map('i', 'jj', '<Esc>')
map('i', 'kk', '<Esc>')
map('i', 'jk', '<Esc>')
map('i', 'kj', '<Esc>')

map('n', ';', ':')

map("n", "J", "mzJ`z")
map("n", "<C-d>", "<C-d>zz")
map("n", "<C-u>", "<C-u>zz")
map("n", "n", "nzzzv")
map("n", "N", "Nzzzv")

map("n", "<space>e", vim.diagnostic.open_float, { noremap = true, silent = true, buffer = bufnr })
