<template>
  <section class="hero min-h-[calc(100vh-5rem)] bg-base-100">
    <div class="hero-content grid w-full max-w-6xl grid-cols-1 gap-10 px-4 py-12 lg:grid-cols-[1.1fr_0.9fr]">
      <div class="space-y-6">
        <div class="badge badge-primary badge-outline">Study Sessions AI</div>
        <h1 class="max-w-3xl text-4xl font-bold leading-tight text-neutral sm:text-5xl">
          Bienvenido a tus sesiones de estudio inteligentes
        </h1>
        <p class="max-w-xl text-base leading-7 text-base-content/70 sm:text-lg">
          Organiza objetivos, mide progreso y mantén el ritmo con una experiencia enfocada para estudiar mejor.
        </p>
        <div class="flex flex-wrap gap-3">
          <a href="#sesiones" class="btn btn-primary">Comenzar</a>
          <a href="#agenda" class="btn btn-outline">Ver agenda</a>
        </div>
      </div>

      <div id="agenda" class="rounded-lg border border-base-300 bg-base-200 p-5 shadow-sm">
        <div class="mb-5 flex items-center justify-between gap-4">
          <div>
            <p class="text-sm font-semibold uppercase tracking-wide text-primary">Hoy</p>
            <h2 class="text-2xl font-bold text-neutral">Bloque de enfoque</h2>
          </div>
          <div class="radial-progress text-primary" style="--value:72; --size:4.5rem; --thickness:0.45rem;" role="progressbar">
            72%
          </div>
        </div>

        <div class="space-y-3">
          <div class="flex items-center justify-between rounded-md bg-base-100 px-4 py-3">
            <span class="font-medium">Repasar apuntes</span>
            <span class="badge badge-success badge-outline">Listo</span>
          </div>
          <div class="flex items-center justify-between rounded-md bg-base-100 px-4 py-3">
            <span class="font-medium">Resolver ejercicios</span>
            <span class="badge badge-primary">En curso</span>
          </div>
          <div class="flex items-center justify-between rounded-md bg-base-100 px-4 py-3">
            <span class="font-medium">Resumen final</span>
            <span class="badge badge-ghost">Pendiente</span>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>