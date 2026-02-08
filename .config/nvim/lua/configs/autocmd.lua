local highlight_group = vim.api.nvim_create_augroup('YankHighlight', { clear = true })

vim.api.nvim_create_autocmd('TextYankPost', {
  group = highlight_group,
  pattern = '*',
  callback = function()
    vim.highlight.on_yank({
      higroup = 'IncSearch', -- The highlight group used for the flash (default is IncSearch)
      timeout = 150,         -- Duration of the highlight in milliseconds
    })
  end,
})


vim.api.nvim_create_autocmd("FileType", {
  pattern = { "lua", "javascript", "typescript", "js", "ts", "tsx", "jsx", "html" },
  callback = function()
    vim.bo.shiftwidth = 2
    vim.bo.tabstop = 2
    vim.bo.expandtab = true
  end,
})
