vim.pack.add {
  { src = 'https://github.com/neovim/nvim-lspconfig' },
}

vim.lsp.enable({ 'gopls', 'lua_ls', 'ts_ls', 'tombi' })

vim.api.nvim_create_autocmd('LspAttach', {
  callback = function(args)
    local opts = { buffer = args.buf }
    local map = vim.keymap.set

    map('n', 'gd', function() vim.lsp.buf.definition() end, {
      buffer = args.buf,
      desc = 'Go to definition (LSP)',
    })

    map('n', 'K', function() vim.lsp.buf.hover() end, {
      buffer = args.buf,
      desc = 'Show hover documentation',
    })

    map('n', '<C-m>', function() vim.lsp.buf.code_action() end, {
      buffer = args.buf,
      desc = 'Code action',
    })

    map('n', '<leader>gr', function() vim.lsp.buf.references() end, {
      buffer = args.buf,
      desc = 'Find references (LSP)',
    })

    map('n', 'rn', function() vim.lsp.buf.rename() end, {
      buffer = args.buf,
      desc = 'Rename symbol',
    })

    map('n', 'gt', function() vim.lsp.buf.type_definition() end, {
      buffer = args.buf,
      desc = 'Go to type definition (LSP)',
    })

    map('n', '<leader>ft', function() vim.lsp.buf.format() end, {
      buffer = args.buf,
      desc = 'Format buffer',
    })

    map('n', '<leader>ff', function() vim.lsp.buf.implementation() end, {
      buffer = args.buf,
      desc = 'Go to implementation (LSP)',
    })

    -- map("n", "<leader>vws", function() vim.lsp.buf.workspace_symbol() end, opts)
    -- map("n", "<leader>vd", function() vim.diagnostic.open_float() end, opts)
    -- map("n", "[d", function() vim.diagnostic.goto_next() end, opts)
    -- map("n", "]d", function() vim.diagnostic.goto_prev() end, opts)
    -- map("i", "<C-h>", function() vim.lsp.buf.signature_help() end, opts)
  end,
})
