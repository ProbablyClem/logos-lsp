import * as path from "path";
import { workspace, ExtensionContext } from "vscode";

import {
  LanguageClient,
  LanguageClientOptions,
  ServerOptions,
  TransportKind,
} from "vscode-languageclient/node";

let client: LanguageClient;

export function activate(context: ExtensionContext) {
  // The server is implemented in go so we have a singe binary to run
  const serverCommand = "logos-lsp";

  // we run the locally build version
  const debugServerCommand = path.join(
    context.extensionPath,
    "./server/logos-lsp"
  );

  // If the extension is launched in debug mode then the debug server options are used
  // Otherwise the run options are used

  const serverOptions: ServerOptions = {
    run: {
      command: serverCommand,
      transport: TransportKind.stdio,
    },
    debug: {
      command: debugServerCommand,
      transport: TransportKind.stdio,
    },
  };

  // Options to control the language client
  const clientOptions: LanguageClientOptions = {
    // Register the server for all documents by default
    documentSelector: [{ scheme: "file", language: "markdown" }],
    synchronize: {
      // Notify the server about file changes to '.clientrc files contained in the workspace
      fileEvents: workspace.createFileSystemWatcher("**/.clientrc"),
    },
  };

  // Create the language client and start the client.
  client = new LanguageClient(
    "logos-lsp",
    "logos lsp",
    serverOptions,
    clientOptions
  );

  // Start the client. This will also launch the server
  client.start();
}

export function deactivate(): Thenable<void> | undefined {
  if (!client) {
    return undefined;
  }
  return client.stop();
}
