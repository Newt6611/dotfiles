return {
  {
    'neovim/nvim-lspconfig',
    ft = { 'go', 'lua', 'typescript', 'sol', 'solidity' },
    config = function()
      vim.lsp.enable({
        'solidity_ls_nomicfoundation',
        'gopls',
        'lua_ls',
        'ts_ls',
      })
    end,

    keys = {
      {
        'gd',
        function() require('snacks').picker.lsp_definitions() end,
        desc = 'Go to definition (LSP)',
      },
      {
        'K',
        vim.lsp.buf.hover,
        desc = 'Show hover documentation',
      },
      {
        '<C-m>',
        vim.lsp.buf.code_action,
        desc = 'Code action',
      },
      {
        'grr',
        function() require('snacks').picker.lsp_references() end,
        desc = 'Find references (LSP)',
      },
      {
        'rn',
        vim.lsp.buf.rename,
        desc = 'Rename symbol',
      },
      {
        'gt',
        function() require('snacks').picker.lsp_type_definitions() end,
        desc = 'Go to type definition (LSP)',
      },
      {
        '<leader>ft',
        vim.lsp.buf.format,
        desc = 'Format buffer',
      },
      {
        '<leader>ff',
        function() require('snacks').picker.lsp_implementations() end,
        desc = 'Go to implementation (LSP)',
      },
    },
  },
}
