{
  description = "Deployer environment with dotfiles";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-25.05";
    flake-utils.url = "github:numtide/flake-utils";
    dotfiles.url = "github:jreslock/dotfiles";
  };

  outputs = { self, nixpkgs, flake-utils, dotfiles }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
        homeCfg = dotfiles.outputs.homeConfigurations.${system};
      in {
        devShells.default = pkgs.mkShell {
          name = "deployer-dev-shell";

          buildInputs = with pkgs; [
            awscli2
            coreutils
            docker
            docker-buildx
            git
            go
            golangci-lint
            goreleaser
            go-task
            opentofu
            pre-commit
            svu
            terraform-docs
            zsh
          ];

          shellHook = ''
            export SHELL=${pkgs.zsh}/bin/zsh
            export GO111MODULE=on

            echo "🧪 Entering Nix shell. Are we in GitHub Actions? (GITHUB_ACTIONS=$GITHUB_ACTIONS)"

            if [ "$GITHUB_ACTIONS" != "true" ]; then
              echo "🔧 Activating home-manager and dotfiles..."
              ${homeCfg.activationPackage}/activate
            else
              echo "⚙️ Skipping dotfile activation in CI context"
            fi
          '';
        };
      }
    );
}
