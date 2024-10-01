export function Dashboard() {
  return (
    <div className="grid min-h-screen w-full grid-cols-1 md:grid-cols-[220px_1fr] lg:grid-cols-[280px_1fr]">
      {/* Sidebar e Header... */}

      <main className="flex flex-1 flex-col gap-4 p-4 sm:p-6 md:gap-6 lg:p-8">
        <div className="flex items-center">
          <h1 className="text-xl font-semibold sm:text-2xl md:text-3xl">Appointments</h1>
        </div>

        <div className="flex flex-1 items-center justify-center rounded-lg border border-dashed shadow-sm">
          <div className="flex flex-col items-center gap-4 text-center">
            {/* Conteúdo a ser adicionado aqui */}
          </div>
        </div>
      </main>
    </div>
  );
}
