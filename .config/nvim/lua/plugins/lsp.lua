return {
    "neovim/nvim-lspconfig",
    opts = {
        inlay_hints = {
            enabled = false,
        },
        servers = {
            ['*'] = {
                keys = {
                    { "rn", "<cmd>lua vim.lsp.buf.rename()<CR>", has = "definition"},
                    { "gd", "<cmd>Trouble lsp_definitions<CR>", has = "definition"},
                    { "<C-m>", "<cmd>lua vim.lsp.buf.code_action()<CR>", has = "definition"},
                    { "<leader>f", "<cmd>lua vim.lsp.buf.implementation()<CR>", has = "definition"},
                    { "gt", "<cmd>lua vim.lsp.buf.type_definition()<CR>", has = "definition"},
                    { "gr", "<cmd>Telescope lsp_references<CR>", has = "definition"},
                    -- { "gr", "<cmd>Trouble lsp_references<CR>", has = "definition"},
                },
            },
        },
    },
}
