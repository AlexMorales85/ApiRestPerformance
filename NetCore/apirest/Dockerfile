FROM mcr.microsoft.com/dotnet/core/sdk:3.0

WORKDIR /usr/src/app

COPY . .

RUN dotnet build

EXPOSE 5000
CMD ["dotnet", "run"]