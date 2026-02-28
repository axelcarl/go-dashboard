import { api } from "@/lib/api-client";
import { type Payment } from "@/types/api";
import { useMutation, useQueryClient } from "@tanstack/react-query";

async function createPayment(payment: Payment): Promise<void> {
  await api.post("/payments", {
    amount: payment.amount,
    sender: payment.sender,
    recipient: payment.recipient,
  });
}

export function useCreatePayment() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: createPayment,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["payments"] });
    },
  });
}
