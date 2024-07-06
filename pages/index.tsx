import Image from "next/image";
import { Inter } from "next/font/google";
import { useCallback, useEffect, useState } from "react";
import { getUsersInQueue, UserInQueue } from "@/utils/api-utils";
import { CommonHead } from "@/components/CommonHead";

const inter = Inter({ subsets: ["latin"] });

export default function Home() {
  const [total, setTotal] = useState(0);
  const [users, setUsers] = useState<UserInQueue[]>([]);

  const refreshQueue = useCallback(() => {
    console.log("Refreshing the queue...");

    getUsersInQueue().then((data) => {
      setTotal(data.total);
      setUsers(data.users);
    });
  }, []);

  useEffect(() => {
    refreshQueue();

    const interval = setInterval(() => {
      refreshQueue();
    }, 60000);

    return () => {
      clearInterval(interval);
    };
  }, [refreshQueue]);

  console.log({ total, users });

  return (
    <>
      <CommonHead />

      <main
        className={`flex min-h-screen flex-col items-center p-24 ${inter.className}`}
      >
        <div className="relative flex place-items-center before:absolute before:h-[300px] before:w-full sm:before:w-[480px] before:-translate-x-1/2 before:rounded-full  before:blur-2xl before:content-[''] after:absolute after:-z-20 after:h-[180px] after:w-full sm:after:w-[240px] after:translate-x-1/3 after:bg-gradient-conic  after:blur-2xl after:content-[''] before:bg-gradient-to-br before:from-transparent before:to-[#EF15BF]/10 after:from-[#EF15BF]/30 after:via-[#EF15BF]/40 before:lg:h-[360px]">
          <Image
            className="relative drop-shadow-[0_0_0.3rem_#ffffff70]"
            src="/queuerrr-logo.svg"
            alt="queuerrr by techygrrrl"
            title="queuerrr by techygrrrl"
            width={322}
            height={131}
            priority
          />
        </div>

        {total > 0 ? (
          <>
            <p className="text-white mt-[80px] mb-[30px]">
              The following users are in the queue
            </p>

            <table className="w-100 sm:w-10/12 lg:w-6/12 border-solid border border-white border-opacity-20">
              <thead>
                <tr className="border-solid border border-white border-opacity-20">
                  <th className="py-2 px-4 bg-slate-900">User</th>
                  <th className="py-2 px-4 bg-slate-900">Note</th>
                </tr>
              </thead>

              <tbody>
                {users.map((user) => (
                  <tr key={user.twitch_user_id}>
                    <td className="py-2 px-4 border-solid border border-white border-opacity-20">
                      <strong>{user.twitch_username}</strong>
                      {/* <span className="text-sm ml-3 opacity-60">{user.twitch_user_id}</span> */}
                    </td>
                    <td className="py-2 px-4 border-solid border border-white border-opacity-20">
                      {user.notes}
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </>
        ) : (
          <p className="text-white mt-[160px]">The queue is currently empty.</p>
        )}

        <p className="text-center text-white mt-[80px] mb-[30px]">
          Available commands
        </p>

        <table className="w-100 sm:w-10/12 lg:w-6/12 border-solid border border-white border-opacity-20">
          <thead>
            <tr className="border-solid border border-white border-opacity-20">
              <th className="py-2 px-4 bg-slate-900">Command</th>
              <th className="py-2 px-4 bg-slate-900">Description</th>
            </tr>
          </thead>

          <tbody>
            <tr>
              <td className="whitespace-nowrap align-top py-2 px-4 border-solid border border-white border-opacity-20">
                <code>!join</code> or <code>!queue join</code>
              </td>
              <td className="py-2 px-4 border-solid border border-white border-opacity-20">
                Join the queue
              </td>
            </tr>

            <tr>
              <td className="whitespace-nowrap align-top py-2 px-4 border-solid border border-white border-opacity-20">
                <code>!leave</code> or <code>!queue leave</code>
              </td>
              <td className="py-2 px-4 border-solid border border-white border-opacity-20">
                Leave the queue
              </td>
            </tr>

            <tr>
              <td className="whitespace-nowrap align-top py-2 px-4 border-solid border border-white border-opacity-20">
                <code>!position</code> or <code>!queue position</code>
              </td>
              <td className="align-top py-2 px-4 border-solid border border-white border-opacity-20">
                See your position in the queue
              </td>
            </tr>

            <tr>
              <td className="whitespace-nowrap align-top py-2 px-4 border-solid border border-white border-opacity-20">
                <code>!info</code> or <code>!queue info</code>
              </td>
              <td className="align-top py-2 px-4 border-solid border border-white border-opacity-20">
                Get a list of everyone in the queue. This page, however, refreshes automatically.
              </td>
            </tr>
          </tbody>
        </table>

        <footer className="mt-20">
          <p className="text-center">
            Want a queuing system? You can{" "}
            <a
              className="text-[#EF15BF] hover:text-[#EBEF15] hover:underline"
              href="https://github.com/techygrrrl/queuerrr"
              target="blank"
            >
              useQueuerrr()
            </a>
          </p>
        </footer>
      </main>
    </>
  );
}
