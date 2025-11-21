return {
    {
        "yanganto/move.vim",
        branch = "sui-move",
        config = function()
            local lspconfig = require("lspconfig")
            local configs = require("lspconfig.configs")
            if not configs.sui_move_analyzer then
                configs.sui_move_analyzer = {
                    default_config = {
                        cmd = { os.getenv("HOME") .. "/.sui/bin/move-analyzer" },
                        filetypes = { "move" },
                        root_dir = lspconfig.util.root_pattern("Move.toml", ".git"),
                    }
                }
            end
            lspconfig.sui_move_analyzer.setup({})
        end
    },
}

