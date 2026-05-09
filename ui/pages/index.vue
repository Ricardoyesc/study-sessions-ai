<template>
  <div data-theme="study" class="min-h-screen bg-base-100 text-base-content">
    <header class="navbar sticky top-0 z-10 border-b border-base-300 bg-base-100/90 px-4 backdrop-blur md:px-8">
      <div class="navbar-start">
        <span class="text-lg font-bold text-neutral">Study Sessions AI</span>
      </div>
      <nav class="navbar-end gap-2">
        <a href="#sesiones" class="btn btn-ghost btn-sm">Sesiones</a>
        <a href="#agenda" class="btn btn-primary btn-sm">Agenda</a>
      </nav>
    </header>

    <main>
      <WelcomeHero />

      <section id="sesiones" class="bg-base-200 py-12">
        <div class="mx-auto grid max-w-6xl gap-6 px-4 md:px-8">
          <div class="flex flex-col gap-5 sm:flex-row sm:items-end sm:justify-between">
            <div>
              <p class="font-semibold uppercase tracking-wide text-primary">Plan de hoy</p>
              <h2 class="text-3xl font-bold text-neutral">Tres pasos para entrar en foco</h2>
            </div>

            <div class="stats stats-vertical shadow-sm sm:stats-horizontal">
              <div class="stat py-3">
                <div class="stat-title">Racha</div>
                <div class="stat-value text-primary">7</div>
              </div>
              <div class="stat py-3">
                <div class="stat-title">Minutos</div>
                <div class="stat-value text-secondary">90</div>
              </div>
            </div>
          </div>

          <div class="grid gap-4 md:grid-cols-3">
            <StudySessionCard
              title="Calentar memoria"
              description="Revisa los conceptos clave y marca lo que necesita una segunda vuelta."
              duration="15 min"
              status="Ligero"
            />
            <StudySessionCard
              title="Practicar activo"
              description="Trabaja ejercicios cortos con pausas medidas para sostener el progreso."
              duration="45 min"
              status="Principal"
            />
            <StudySessionCard
              title="Cerrar con resumen"
              description="Convierte lo aprendido en notas claras para la siguiente sesión."
              duration="30 min"
              status="Refuerzo"
            />
          </div>
        </div>
      </section>
    </main>
  </div>
</template>