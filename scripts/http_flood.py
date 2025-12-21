import asyncio
import aiohttp
import time
from collections import defaultdict

async def send_request(session, url, data, semaphore, results):
    async with semaphore:
        try:
            async with session.post(url, json=data) as response:
                status = response.status
                results[status] += 1
                return status
        except Exception as e:
            results['errors'] += 1
            return str(e)

async def main():
    url = "http://localhost/posts"
    data = {
        "title": "Load Test Post",
        "content": "This is a load test content.",
        "author": "Pedro"
    }
    
    requests_per_second = 100000
    duration_seconds = 20
    total_requests = requests_per_second * duration_seconds
    
    results = defaultdict(int)
    semaphore = asyncio.Semaphore(1000)
    
    connector = aiohttp.TCPConnector(limit=1000, limit_per_host=1000)
    timeout = aiohttp.ClientTimeout(total=30)
    
    async with aiohttp.ClientSession(
        connector=connector, 
        timeout=timeout
    ) as session:
        start_time = time.time()
        
        for second in range(duration_seconds):
            second_start = time.time()
            
            tasks = [
                send_request(session, url, data, semaphore, results)
                for _ in range(requests_per_second)
            ]
            
            await asyncio.gather(*tasks)
            
            second_elapsed = time.time() - second_start
            total_completed = sum(results.values())
            
            print(
                f"[{second + 1}/{duration_seconds}] "
                f"Completed: {total_completed}/{total_requests} | "
                f"Second duration: {second_elapsed:.2f}s | "
                f"2xx: {results[200] + results[201]} | "
                f"Errors: {results['errors']}"
            )
            
            elapsed = time.time() - start_time
            expected_elapsed = second + 1
            sleep_time = expected_elapsed - elapsed
            if sleep_time > 0:
                await asyncio.sleep(sleep_time)
        
        end_time = time.time()
        
        print("\n=== SUMMARY ===")
        print(f"Total time: {end_time - start_time:.2f} seconds")
        print(f"Total requests: {total_requests}")
        print(f"RPS: {total_requests / (end_time - start_time):.2f}")
        # print("\nStatus codes:")
        # for status, count in sorted(results.items()):
        #     if status != 'errors':
        #         print(f"  {status}: {count}")
        # if results['errors']:
        #     print(f"  Errors: {results['errors']}")

if __name__ == "__main__":
    asyncio.run(main())
