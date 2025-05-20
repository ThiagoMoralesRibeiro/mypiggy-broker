# mypiggy-broker

🏠 go-home-broker
Este serviço Go simula um home broker, processando ordens de compra e venda de ações em tempo real. Ele faz parte do projeto mypiggy-broker e atua como um componente central para o roteamento de ordens.

🚀 Pré-requisitos
Go 1.20 ou superior.

Git instalado.

(Opcional) Docker, caso deseje executar via contêiner.
WIRED
+1
Make a README
+1

📦 Clonando o repositório
bash
Copiar
Editar
git clone https://github.com/ThiagoMoralesRibeiro/mypiggy-broker.git
cd mypiggy-broker/go-home-broker
🔨 Build do serviço
bash
Copiar
Editar
go build -o home-broker main.go
Isso gerará um executável chamado home-broker no diretório atual.

▶️ Executando o serviço
bash
Copiar
Editar
./home-broker
O serviço será iniciado e escutará na porta padrão (ex: :8080).

🐳 Executando com Docker (opcional)
bash
Copiar
Editar
docker build -t go-home-broker .
docker run -p 8080:8080 go-home-broker
📁 Estrutura do projeto
go
Copiar
Editar

go-home-broker/
├── main.go
├── go.mod
├── go.sum
├── internal/
│   ├── domain/
│   ├── handler/
│   └── service/
└── README.md
🧪 Testes
Para executar os testes automatizados:

bash
Copiar
Editar
go test ./...
📄 Licença
Este projeto está licenciado sob a licença MIT.

