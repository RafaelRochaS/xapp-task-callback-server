import json
import statistics
from datetime import datetime
from collections import Counter

def load_jsonl(path):
    results = []
    with open(path, "r") as f:
        for line in f:
            line = line.strip()
            if not line:
                continue
            try:
                results.append(json.loads(line))
            except json.JSONDecodeError:
                print(f"Linha inválida ignorada: {line}")
    return results

def percentile(data, p):
    if not data:
        return None
    k = (len(data)-1) * (p/100)
    f = int(k)
    c = min(f+1, len(data)-1)
    if f == c:
        return data[f]
    return data[f] + (data[c] - data[f]) * (k - f)

def main():
    path = "results.jsonl"

    results = load_jsonl(path)
    if not results:
        print("Nenhum dado encontrado.")
        return

    # Extrair valores
    latencies = []
    durations = []
    workloads = []
    start_times = []
    end_times = []
    locations = []

    for r in results:
        try:
            durations.append(float(r["duration"]))           # duração da tarefa
            workloads.append(int(r["workloadSize"]))          # tamanho workload
            start_times.append(float(r["createdAt"]))
            end_times.append(float(r["timestamp"]))
            latencies.append(float(r["timestamp"]) - float(r["createdAt"]))
            locations.append(r.get("executionSite", "unknown"))
        except Exception as e:
            print("Erro processando linha:", e)

    total_runtime = max(end_times) - min(start_times)
    throughput = len(results) / total_runtime if total_runtime > 0 else 0

    # Estatísticas básicas
    print("\n=== MÉTRICAS DA SIMULAÇÃO ===\n")

    print(f"Total de execuções: {len(results)}")
    print(f"Tempo total da simulação: {total_runtime:.3f} s")
    print(f"Throughput médio: {throughput:.3f} tasks/s")

    print("\n--- Workload ---")
    print(f"Média workload: {statistics.mean(workloads):.2f}")
    print(f"Desvio padrão: {statistics.pstdev(workloads):.2f}")
    print(f"Máximo: {max(workloads)}  Mínimo: {min(workloads)}")

    print("\n--- Latência ---")
    print(f"Latência média: {statistics.mean(latencies):.4f} s")
    print(f"P50: {percentile(latencies, 50):.4f}")
    print(f"P90: {percentile(latencies, 90):.4f}")
    print(f"P95: {percentile(latencies, 95):.4f}")
    print(f"P99: {percentile(latencies, 99):.4f}")

    print("\n--- Duração da Execução ---")
    print(f"Duração média: {statistics.mean(durations):.4f} s")
    print(f"P90 duração: {percentile(durations, 90):.4f}")

    print("\n--- Local de Execução ---")
    count = Counter(locations)
    for loc, qty in count.items():
        print(f"{loc}: {qty} execuções")

    print("\n=== FIM ===")

if __name__ == "__main__":
    main()
