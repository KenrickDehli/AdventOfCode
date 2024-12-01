FROM mcr.microsoft.com/dotnet/runtime:6.0 AS base
WORKDIR /app

FROM mcr.microsoft.com/dotnet/sdk:6.0 AS build
WORKDIR /src
COPY ["AoCDayOne/AoCDayOne.csproj", "AoCDayOne/"]
RUN dotnet restore "AoCDayOne/AoCDayOne.csproj"
COPY . .
WORKDIR "/src/AoCDayOne"
RUN dotnet build "AoCDayOne.csproj" -c Release -o /app/build

FROM build AS publish
RUN dotnet publish "AoCDayOne.csproj" -c Release -o /app/publish /p:UseAppHost=false

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "AoCDayOne.dll"]
