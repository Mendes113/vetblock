"use client"

import * as React from "react"
import { TrendingUp } from "lucide-react"
import { Label, Pie, PieChart, Legend } from "recharts"

import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart"

// Descrição do gráfico
export const description = "A donut chart showing consultations with animals."

// Dados atualizados para representar consultas com cães e gatos
const chartData = [
  { animal: "Cachorros", consultas: 5, fill: "var(--color-dog)" },
  { animal: "Gatos", consultas: 10, fill: "var(--color-cat)" },
]

const chartConfig = {
  consultas: {
    label: "Consultas",
  },
  dog: {
    label: "Cachorros",
    color: "hsl(var(--chart-1))", // Cor personalizada para cachorros
  },
  cat: {
    label: "Gatos",
    color: "hsl(var(--chart-2))", // Cor personalizada para gatos
  },
} satisfies ChartConfig

export function ChartPie() {
  const totalConsultas = React.useMemo(() => {
    return chartData.reduce((acc, curr) => acc + curr.consultas, 0)
  }, [])

  return (
    <Card className="flex flex-col">
      <CardHeader className="items-center pb-0">
        <CardTitle>Consultas da Semana</CardTitle>
        <CardDescription>Outubro 2024</CardDescription>
      </CardHeader>
      <CardContent className="flex-1 pb-0">
        <ChartContainer
          config={chartConfig}
          className="mx-auto aspect-square max-h-[250px]"
        >
          <PieChart>
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent hideLabel />}
            />
            <Pie
              data={chartData}
              dataKey="consultas"
              nameKey="animal"
              innerRadius={60}
              strokeWidth={5}
            >
              <Label
                content={({ viewBox }) => {
                  if (viewBox && "cx" in viewBox && "cy" in viewBox) {
                    return (
                      <text
                        x={viewBox.cx}
                        y={viewBox.cy}
                        textAnchor="middle"
                        dominantBaseline="middle"
                      >
                        <tspan
                          x={viewBox.cx}
                          y={viewBox.cy}
                          className="fill-foreground text-3xl font-bold"
                        >
                          {totalConsultas.toLocaleString()}
                        </tspan>
                        <tspan
                          x={viewBox.cx}
                          y={(viewBox.cy || 0) + 24}
                          className="fill-muted-foreground"
                        >
                          Consultas
                        </tspan>
                      </text>
                    )
                  }
                }}
              />
            </Pie>
            {/* Adiciona a legenda ao gráfico */}
            <Legend
              layout="horizontal"
              align="center"
              verticalAlign="bottom"
              payload={[
                { value: "Cachorros", type: "circle", color: "var(--color-dog)" },
                { value: "Gatos", type: "circle", color: "var(--color-cat)" },
              ]}
            />
          </PieChart>
        </ChartContainer>
      </CardContent>
      <CardFooter className="flex-col gap-2 text-sm">
        <div className="flex items-center gap-2 font-medium leading-none">
          Consultas aumentando 5.2% esta semana <TrendingUp className="h-4 w-4" />
        </div>
        <div className="leading-none text-muted-foreground">
          Exibindo consultas de animais da semana atual
        </div>
      </CardFooter>
    </Card>
  )
}
