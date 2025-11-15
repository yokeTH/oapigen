{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = {
    self,
    nixpkgs,
    flake-utils,
    ...
  }:
    flake-utils.lib.eachDefaultSystem (
      system: let
        goVersion = "1.25";

        goAttr = "go_" + builtins.replaceStrings ["."] ["_"] goVersion;

        pkgs = import nixpkgs {
          inherit system;
          overlays = [
            (final: prev: {
              go = prev.${goAttr};
            })
          ];
        };

        app = pkgs.buildGoModule {
          pname = "oapigen";
          version = "0.1.0";

          src = ./.;

          vendorHash = pkgs.lib.fakeSha256;

          nativeBuildInputs = [];

          preBuild = ''
          '';

          meta = with pkgs.lib; {
            description = "Code to open API specification without describe";
            license = licenses.bsd3;
          };
        };
      in {
        packages = {
          default = app;
          app = app;
        };

        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            gopls
            golangci-lint
            pre-commit
          ];

          shellHook = ''
          '';
        };
      }
    );
}
