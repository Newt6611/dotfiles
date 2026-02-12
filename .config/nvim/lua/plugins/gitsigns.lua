vim.pack.add({
  { src = 'https://github.com/lewis6991/gitsigns.nvim' }
})

require('gitsigns').setup({
  -- current_line_blame = true,
  -- current_line_blame_opts = {
  --   virt_text = true,
  --   virt_text_pos = 'eol', -- 'eol' | 'overlay' | 'right_align'
  --   delay = 0,
  --   ignore_whitespace = false,
  --   virt_text_priority = 100,
  --   use_focus = true,
  -- },
})

local map = vim.keymap.set

map({ 'n', 'v' }, '<leader>r', ':Gitsigns reset_hunk<CR>', { desc = 'Reset Hunk' })
map('n', '<leader>b', '<cmd>Gitsigns blame_line<CR>', { desc = 'Reset Hunk' })
