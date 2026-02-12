local map = vim.keymap.set

map('i', 'jj', '<Esc>')
map('i', 'kk', '<Esc>')
map('i', 'jk', '<Esc>')
map({ 'n', 'v' }, ';', ':')

map('n', 'J', 'mzJ`z')
map('n', '<C-d>', '<C-d>zz')
map('n', '<C-u>', '<C-u>zz')
map('n', 'n', 'nzzzv')
map('n', 'N', 'Nzzzv')

map('n', '<space>e', vim.diagnostic.open_float, { noremap = true, silent = true, buffer = bufnr })

map('n', 'zj', ':cnext<CR>', { noremap = true, silent = true })
map('n', 'zk', ':cprev<CR>', { noremap = true, silent = true })
map('n', 'zo', ':copen<CR>', { noremap = true, silent = true })

map('n', '<leader>xx', function()
  vim.diagnostic.setqflist()
end, { desc = 'Diagnostics (Quickfix)' })

map('n', '<leader>em', function()
  -- Open a new empty buffer
  vim.cmd('new')

  -- Set buffer type to 'nofile' so it won't be associated with a file
  vim.bo.buftype = 'nofile'

  -- Hide the buffer when abandoned instead of closing it
  vim.bo.bufhidden = 'hide'

  -- Disable swapfile for this buffer
  vim.bo.swapfile = false

  -- Fill the buffer with the output of :messages, split by line
  vim.api.nvim_buf_set_lines(0, 0, -1, false, vim.split(vim.fn.execute('messages'), '\n'))

  -- Move the cursor to the last line of the buffer
  vim.cmd('normal! G')
end, { desc = 'Open message history' })
